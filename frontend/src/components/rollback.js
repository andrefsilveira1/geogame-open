import { Button, Input } from 'antd';
import React, { useState } from 'react';
import axios from 'axios';
const apiURL = process.env.REACT_APP_API_URL;

const Rollback = ( { setPlayed, setTracker, setCall, callNext, points } ) => {
    const [name, setName] = useState('');

    const handleName = (e) => {
        setName(e.target.value);
    };

    const Reset = async () => {
        setPlayed({Ids: []});
        setTracker(0);
        setCall(!callNext);
        const score = points.points
        if(name)
            await axios.post(`${apiURL}/user/score`, {name, score}).catch();

    };

    return (
        <>
            <h2>Ops, it seems that you played a lot...</h2>
            <p>Sorry, we do not have more countries to play. If you want, you can send your score to the billboard</p>
            <Input placeholder="Your name here"  value={name} onChange={handleName}/>
            <p>Also, you can rollback and play again!</p>
            <Button className="rollback-button" onClick={Reset}>Rollback</Button>
        </>
    )
}

export default Rollback