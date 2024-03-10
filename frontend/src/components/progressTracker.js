import React, { useEffect } from 'react';
import { Steps, Button, Flex } from 'antd';
import "./progress.css"
const Progresstracker = ({ tracker, country, setTracker, tipTracker }) => {


    useEffect(() => {
        setTracker(0);
    }, [country])

    const nextTip = () => {
        setTracker(tipTracker + 1);
        tracker(tipTracker + 1);
    }
    let countryItems
    if (country) {
        countryItems = country?.tips?.map((tip, index) => ({
            description: tipTracker >= index + 1 ? tip.text : '',
            subTitle: `Score bonus ${(50 / (1 + index)).toFixed(0)}%`,
        }));
    }


    return (

        <>
            <Flex vertical align='center' gap="small" style={{ width: "100%" }}>
                <Steps
                    current={tipTracker}
                    className="progress-tracker"
                    items={countryItems}
                />
                <Flex vertical gap="small" style={{ width: '25%' }}>
                    <Button type="primary" className="tip-button" onClick={nextTip}>
                        Use Hint
                    </Button>
                </Flex>
            </Flex>
        </>
    )
};
export default Progresstracker;