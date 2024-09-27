import React, { createContext, useState, useContext } from 'react';

const UrlContext = createContext();

export const UrlProvider = ({ children }) => {
  const [shortenedUrls, setShortenedUrls] = useState([]);

  //setting the intial state
  const setUrls = (urls) => {
    setShortenedUrls(urls); 
  };

  //adding new url
  const addShortenedUrl = (newUrl) => {
    setShortenedUrls((initialUrls) => [...initialUrls, newUrl]);
  };

  return (
    <UrlContext.Provider value={{ shortenedUrls, addShortenedUrl, setUrls }}>
      {children}
    </UrlContext.Provider>
  );
};

export const useUrlContext = () => {
  return useContext(UrlContext);
};
