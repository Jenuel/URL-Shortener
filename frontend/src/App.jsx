import { useState } from 'react'
import './App.css'
import AdminPage from './pages/AdminPage'
import { UrlProvider } from './context/UrlContext';

function App() {
 

  return (
    <UrlProvider>
       <AdminPage/>
    </UrlProvider>
  )
}

export default App
