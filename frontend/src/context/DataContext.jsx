import React, { createContext, useState, useContext } from 'react';

const DataContext = createContext();

export function useData() {
  return useContext(DataContext);
}

export function DataProvider({ children }) {
  const [data, setData] = useState([
    // initial data
  ]);

  const addData = (newEntry) => {
    setData(prevData => [...prevData, newEntry]);
  };

  return (
    <DataContext.Provider value={{ data, addData }}>
      {children}
    </DataContext.Provider>
  );
}
