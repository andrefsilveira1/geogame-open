import React from 'react';
import { Modal, Timeline } from 'antd';
const HowPlayCard = ({ isModalOpen, setHowPlay }) => {

    const handleCancel = () => {
        setHowPlay(false);
    };

    return (

        <>
            <Modal title="How to play" open={isModalOpen} onOk={handleCancel} onCancel={handleCancel}>
                <Timeline style={{marginTop: '35px'}}
                    
                    items={[
                        {
                            children: 'A random country will be given',
                            color: 'green'
                        },
                        {
                            children: 'You will try to figure out this country by your geolocalization',
                            color: 'green'

                        },
                        {
                            children: 'Use hints to help you',
                            color: 'green'

                        },
                        {
                            children: 'Less hints, more hint bonus. More hints, less hint bonus',
                            color: 'green'

                        },
                        {
                            children: 'Move the map around the world and click where you think it is the right place. Finally, click in "Send location"',
                            color: 'green'

                        },
                        {
                            children: 'Pay attention: the right location it will be always at the capital of the given country, and you have only one attempt per country',
                            color: 'green'

                        },
                        {
                            children: 'Your final score will be calculated by the distance from your pick compared to the right point and how many tips you used',
                            color: 'green'
                        },
                    ]}
                />
            </Modal>
        </>
    );
};
export default HowPlayCard;