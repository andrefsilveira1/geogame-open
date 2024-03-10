import React, { useEffect } from 'react';
import { Modal, Typography } from 'antd';
const { Text } = Typography;
const ResultCard = ({ isModalOpen, result, setModal, tip, setTip, setPoints, setCall, callNext}) => {
  const handleCancel = () => {
    setModal(false);
    setTip(0);
    setCall(!callNext);
  };
  const score = Number(result?.Score)
  const bonus = Number(tip) === 0 ? (score * 0.60).toFixed(2) : (score * (0.50 / Number(tip)).toFixed(2))
  const final = Number(tip) === 0 ? (score + (score * 0.60)).toFixed(2) : (score + (score * (0.50 / Number(tip)))).toFixed(2)
  const percentage = Number(tip) === 0  ? 0.60 : (0.50 / Number(tip)).toFixed(2)
  
  useEffect(() => {
    const score = Number(result?.Score);
    const final = Number(tip) === 0 ? (score + (score * 0.60)) : (score + (score * (0.50 / Number(tip))));

    setPoints((prevData) => ({
      ...prevData,
      points: isNaN(final) ? 0 : prevData.points + final,
    }));

  }, [result]);



  return (

    <>
      <Modal title="And your result is..." open={isModalOpen} onOk={handleCancel} onCancel={handleCancel}>
        <h2>{result?.Message}</h2>
        <p>Score: {result?.Score}</p>
        <p>Hints used: {Number(tip)} </p>
        <p>Bonus: {bonus} points - <Text type='secondary'> {percentage * 100}%</Text></p>
        <p>Final score: {final}</p>
        <p>Distance: {result?.Distance?.toFixed(2) ?? 'N/A'} Km</p>
        <p>Country: {result?.CountryName}</p>
      </Modal>
    </>
  );
};
export default ResultCard;