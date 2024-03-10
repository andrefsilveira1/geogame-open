import React, { useState, useEffect } from 'react';
import axios from "axios";
import "./main.css";
import logo from "./logo.png";
import { Breadcrumb, Layout, Menu, theme, Button, Flex, Switch, Typography } from 'antd';
import { BugOutlined } from '@ant-design/icons'
import Map from '../components/map';
import Progresstracker from '../components/progressTracker';
import ResultCard from '../components/card';
import HowPlayCard from '../components/howtoplay';
import Info from '../components/info';
import Rollback from '../components/rollback';
import Bug from '../components/bug';
const { Header, Content, Sider } = Layout;
const { Text } = Typography;

const apiURL = process.env.REACT_APP_API_URL

const Main = () => {
    const {
        token: { colorBgContainer, borderRadiusLG },
    } = theme.useToken();

    const [userAnswer, setAnswer] = useState();
    const [country, setCountry] = useState();
    const [tip, setTip] = useState(0);
    const [isModalOpen, setModal] = useState(false);
    const [result, setResult] = useState(false);
    const [playedCountries, setPlayed] = useState({ Ids: [] });
    const [howPlayOpen, setHowPlay] = useState(false);
    const [bug, setBug] = useState(false);
    const [rounds, setRounds] = useState(0);
    const [points, setPoints] = useState({ points: 0 });
    const [tipTracker, setTracker] = useState(0);
    const [callNext, setCall] = useState(false);
    const [players, setPlayers] = useState([]);

    useEffect(() => {
        const getPlayers = async () => {
            try {
                const topPlayers = await axios.get(`${apiURL}/users/score`).catch();
                const top10Scores = topPlayers.data
                    .sort((a, b) => b.score - a.score)
                    .slice(0, 10).filter(function (element) {
                        return element !== "";
                    });

                const menuItems = top10Scores.map((item, index) => ({
                    key: `${index + 1}`,
                    label: `${index === 0 ? '⭐ ' : ''}${item.name} - ${item.score.toFixed(0)} points`,
                }));

                setPlayers(menuItems);
            } catch (err) {
                console.log("ERROR:", err)
            }
        }

        getPlayers();
    }, []);


    const sendPosition = async () => {
        try {
            const data = {
                latitude: userAnswer.lat,
                longitude: userAnswer.lng,
                countryId: country.id,
                time: 0,
            }

            setPlayed((prevPlayedCountries) => ({
                ...prevPlayedCountries,
                Ids: [...prevPlayedCountries.Ids, data.countryId],
            }));
            const response = await axios.post(`${apiURL}/answer`, data).catch(error => console.log("ERROR: Failed to get random country", error));
            setResult(response.data);
            setModal(true);

        } catch (err) {
            console.log("ERROR:", err);
        }
    }

    useEffect(() => {
        const getRandomCountry = async () => {
            try {
                const data = {
                    played: playedCountries
                };

                const response = await axios.post(`${apiURL}/random/country`, data).catch(error => console.log("ERROR: Failed to get random country", error));
                setCountry(response.data);
                setRounds(rounds + 1);
            } catch (err) {
                console.log("ERROR", err);
            }
        };


        getRandomCountry();
    }, [callNext]);

    const handleMenuClick = (key) => {
        if (key === 'how') {
            setHowPlay(true);
        } else if (key === 'bug') {
            setBug(true);
        }
    }

    return (
        <Layout>
            <Header
                style={{
                    display: 'flex',
                    alignItems: 'center',
                    background: "#262626",
                }}
            >
                <img src={logo} alt="Logo" style={{ maxHeight: '100%', maxWidth: '100%', marginRight: '25px' }} />
                <Menu
                    theme="dark"
                    mode="horizontal"
                    defaultSelectedKeys={['play']}
                    onClick={({ key }) => handleMenuClick(key)}
                    style={{
                        background: "#262626",
                    }}
                >
                    <Menu.Item key="play">Play</Menu.Item>
                    <Menu.Item key="how">How to play</Menu.Item>
                    <Menu.Item key="bug"> <BugOutlined /> Found a bug ? </Menu.Item>
                </Menu>
                <p style={{ color: 'white' }}>EN-US</p> <Switch style={{ marginLeft: '7px', marginRight: '7px' }} disabled /><p style={{ color: 'white' }}>PT-BR  <Text type="warning">(Soon)</Text></p>
            </Header>
            <Layout>
                <Sider
                    width={200}
                    collapsedWidth={0}
                    breakpoint="md"
                    style={{
                        background: '#262626',
                    }}
                >
                    <h3 style={{ color: 'white', marginLeft: '16px', marginBottom: '16px' }}>Top 10 players</h3>
                    <Menu
                        mode="inline"
                        defaultSelectedKeys={['1']}
                        defaultOpenKeys={['sub1']}
                        className="custom-menu"
                        style={{
                            borderRight: 0,
                            background: '#262626',
                            height: '100vh',
                        }}
                        items={players}
                    ></Menu>

                </Sider>
                <Layout
                    style={{
                        padding: '0 10px 24px',
                        color: '#262626',
                    }}
                >
                    <Breadcrumb
                        style={{
                            margin: '16px 0',
                        }}
                    >
                        <Breadcrumb.Item>Home</Breadcrumb.Item>
                    </Breadcrumb>
                    <Content
                        style={{
                            padding: 24,
                            margin: 0,
                            minHeight: 'calc(100vh - 131px)',
                            background: colorBgContainer,
                            borderRadius: borderRadiusLG,
                        }}
                    >
                        {country !== null ? (
                            <>
                                <Progresstracker tracker={setTip} country={country} setPlayed={setPlayed} setTracker={setTracker} tipTracker={tipTracker} />
                                <Flex vertical justify="center">
                                    <Map setAnswer={setAnswer} />
                                    <Button type="primary" className="button-loc" onClick={sendPosition}>Send Location</Button>
                                    <Info points={points} rounds={rounds} />
                                </Flex>
                            </>
                        ) : (
                            <>
                                <Rollback setTracker={setTracker} setPlayed={setPlayed} callNext={callNext} setCall={setCall} points={points} />
                                <Info points={points} rounds={rounds} />
                            </>
                        )}
                        <ResultCard
                            isModalOpen={isModalOpen}
                            setModal={setModal}
                            result={result}
                            tip={tip}
                            setTip={setTip}
                            setPoints={setPoints}
                            setCall={setCall}
                            callNext={callNext}
                        />
                        <HowPlayCard isModalOpen={howPlayOpen} setHowPlay={setHowPlay} />
                        <Bug isModalOpen={bug} setHow={setBug} />
                    </Content>
                </Layout>
            </Layout>
        </Layout>
    );
};
export default Main;
