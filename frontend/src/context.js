import React, { createContext, useContext, useState } from 'react';

const LanguageContext = createContext({ language: 'pt-br', toggleLanguage: () => {} });

const LanguageProvider = ({ children }) => {
  const [language, setLanguage] = useState('pt-br');

  const toggleLanguage = () => {
    setLanguage((prevLanguage) => (prevLanguage === 'pt-br' ? 'en-us' : 'pt-br'));
  };

  return (
    <LanguageContext.Provider value={{ language, toggleLanguage }}>
      {children}
    </LanguageContext.Provider>
  );
};

const useLanguage = () => {
  const context = useContext(LanguageContext);
  if (!context) {
    throw new Error('useLanguage must be used within a LanguageProvider');
  }
  return context;
};

export { LanguageProvider, useLanguage };
