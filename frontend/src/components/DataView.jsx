import React, { useEffect, useState } from 'react'
import axios from 'axios';

function DataView() {
  useEffect(() => {
    const fetchUrls = async () => {
      try {
        const response = await axios.get('http://localhost:1323/links');
        setData(response.data);
      } catch (error) {
        console.error('Error fetching portfolios:', error);
      } finally {
        setLoading(false);
      }
    }
  
    fetchUrls();
  })

  const [data, setData] = useState([])
  const [loading, setLoading] = useState(true)

  if (loading) {
    return <div>Loading...</div>;
  }
  
  return (
    <div>
    <h1>Data View</h1>
    <table>
      <thead>
        <tr>
          <th>Original URL</th>
          <th>Short Code</th>
          <th>Clicks</th>
        </tr>
      </thead>
      <tbody>
        {data.map(item => (
          <tr key={item.id}>
            <td><a href={item.original} target="_blank" rel="noopener noreferrer">{item.original}</a></td>
            <td>{item.short}</td>
            <td>{item.click}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>
  )
}

export default DataView