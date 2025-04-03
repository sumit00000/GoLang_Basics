const { Order_Model } = require('../../../models/order');
const { ROBOT_PATH_RESULT, DOCK_CYCLE_STATUS } = require('../../../config/mqtt-topic');

async function updateOrderStatus(orderId, statusData, stationType) {
    try {
        // Validation of statusData
        if (!statusData) {
            throw new Error(`Invalid status data for orderId ${orderId}`);
        }

        // Helper function to safely extract the Boolean or null value
        function getStatus(property) {
            if (property === null || property === undefined) {
                return null; // Return null if the property is null or undefined
            }
            return !!property; // Cast property to Boolean: null or undefined will return false, otherwise true
        }

        // Dynamically set the station status field (source or destination)
        const statusField = stationType === "source" ? "sourceStatus" : "destinationStatus";

        const updateData = {
            $set: { 
                [`orderStatus.${statusField}.followPath`]: getStatus(statusData.followPath.success),
                [`orderStatus.${statusField}.dockForConveyor`]: getStatus(statusData.dockForConveyor.success),
                [`orderStatus.${statusField}.palletClockwise`]: getStatus(statusData.palletClockwise.success),
                [`orderStatus.${statusField}.palletAnticlockwise`]: getStatus(statusData.palletAnticlockwise.success),
                [`orderStatus.${statusField}.undockRobot`]: getStatus(statusData.undockRobot.success),
                [`orderStatus.${statusField}.result`]: getStatus(statusData.finalResult.success),
                [`orderStatus.${statusField}.lastUpdated`]: new Date() // Update time for source or destination
            }
        };

        const updatedOrder = await Order_Model.findOneAndUpdate(
            { order_id: orderId },
            updateData,
            { new: true }
        );

        if (!updatedOrder) {
            throw new Error(`Order with order_id ${orderId} not found`);
        }

        console.log(`Order ${orderId} ${stationType} status updated successfully.`);
        return updatedOrder;
    } catch (error) {
        console.error(`Failed to update order ${orderId} ${stationType} status: ${error.message}`);
        throw error;
    }
}



async function monitorPathStatus(client, amrId, orderId, stationType) {
    return new Promise((resolve, reject) => {
        const topic = `${ROBOT_PATH_RESULT}${amrId}`;

        client.subscribe(topic, (err) => {
            if (err) {
                return reject(new Error(`Failed to subscribe to topic ${topic}: ${err.message}`));
            }
            console.log(`Successfully subscribed to topic: ${topic}`);
        });

        const messageHandler = async (incomingTopic, message) => {
            try {
                if (incomingTopic.startsWith(ROBOT_PATH_RESULT)) {
                    const data = JSON.parse(message.toString());
                    console.log(`Current Order Result: ${topic}, Parsed Data:`, data);

                    // await updateOrderStatus(orderId, { followPath: data.followPath }, stationType);

                    if (data.result !== null) {
                        if (data.result) {
                            console.log(`Robot ${amrId} successfully completed the path nodes`);
                            resolve(true);
                        } else {
                            console.log(`Robot ${amrId} failed to complete the path nodes`);
                            resolve(false);
                        }
                    } else {
                        console.error(`Result of Robot ${amrId} is null, handling error`);
                        resolve(false);
                    }

                    // Clean up after resolution
                    client.unsubscribe(topic, (err) => {
                        if (err) {
                            console.error(`Failed to unsubscribe from topic ${topic}: ${err.message}`);
                        } else {
                            console.log(`Successfully unsubscribed from topic: ${topic}`);
                        }
                    });
                    client.removeListener('message', messageHandler);
                }
            } catch (error) {
                console.error("Error parsing MQTT message:", error.message);
                resolve(false);
            }
        };

        client.on('message', messageHandler);
    });
}

async function monitorDockStatus(client, amrId, orderId, stationType) {
    return new Promise((resolve, reject) => {
        const topic = `${DOCK_CYCLE_STATUS}${amrId}`;

        client.subscribe(topic, (err) => {
            if (err) {
                reject(new Error(`Failed to subscribe to topic ${topic}: ${err.message}`));
            } else {
                console.log(`Successfully subscribed to topic: ${topic}`);
            }
        });

        client.on('message', async (incomingTopic, message) => {
            try {
                if (incomingTopic.startsWith(DOCK_CYCLE_STATUS)) {
                    const data = JSON.parse(message.toString());
                    console.log(`Current Order Result: ${DOCK_CYCLE_STATUS}${amrId}, Parsed Data:`, data);

                    // Now we directly use the stationType passed as argument
                    // await updateOrderStatus(orderId, {
                    //     followPath: data.followPath,
                    //     dockForConveyor: data.dockForConveyor,
                    //     palletClockwise: data.palletClockwise,
                    //     palletAnticlockwise: data.palletAnticlockwise,
                    //     undockRobot: data.undockRobot,
                    //     finalResult: data.finalResult,
                    // }, stationType); // Use the passed stationType directly

                    if (data.finalResult && data.finalResult.success !== null) {
                        if (data.finalResult.success) {
                            console.log(`Robot ${amrId} reached the ${stationType} station successfully.`);
                            resolve(true);
                        } else {
                            console.error(`Robot ${amrId} failed to reach the ${stationType} station: ${data.finalResult.message}`);
                            reject(new Error(`Robot failed to reach ${stationType} station: ${data.finalResult.message}`));
                        }
                    }
                }
            } catch (error) {
                console.error("Error parsing MQTT message:", error.message);
            }
        });
    });
}


module.exports = { monitorPathStatus, monitorDockStatus };


