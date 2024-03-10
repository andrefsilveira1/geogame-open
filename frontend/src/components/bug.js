import React from 'react';
import { Modal, Timeline, Typography } from 'antd';
import { GithubOutlined, MailOutlined } from '@ant-design/icons';

const { Paragraph, Link } = Typography;

const Bug = ({ isModalOpen, setHow }) => {
  const handleCancel = () => {
    setHow(false);
  };

  return (
    <>
      <Modal title="Found a bug?" open={isModalOpen} onOk={handleCancel} onCancel={handleCancel}>
        <Paragraph>
          If you found a bug, please do not hesitate to reach me out. You can send me an email
          or create an issue on GitHub. Describe the bug and if you can, provide an image.
        </Paragraph>
        <Paragraph>
            Email: {' '}
          <Link href="mailto:freitasandre38@gmail.com" target="_blank" rel="noopener noreferrer">
            <MailOutlined /> freitasandre38@gmail.com
          </Link>{' '}
        </Paragraph>
      </Modal>
    </>
  );
};

export default Bug;
