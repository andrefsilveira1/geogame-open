import React from 'react';
import { Flex, Statistic, Card } from 'antd';

const Info = ({ points, rounds}) => {
    rounds = rounds - 1;
    return (
        <Flex justify='space-around' align='center'>
            <Card bordered={true}>
                <Statistic
                    title="Rounds played"
                    value={rounds}
                    valueStyle={{
                        color: '#3f8600',
                    }}
                    suffix={"rounds"}
                />
            </Card>
            <Card bordered={true}>
                <Statistic
                    title="Total score (available)"
                    value={rounds * 8}
                    valueStyle={{
                        color: '#3f8600',
                    }}
                    suffix="points"
                />
            </Card>
            <div>
                <Statistic
                title="Score acquired"
                value={(points.points).toFixed(2)}
                valueStyle={{
                    color: '#3f8600',
                }}
                suffix="points"
            />
            </div>
        </Flex>
    );
};

export default Info;
