import React, { useState } from 'react'

function AdminPage() {
  const [url, setUrl] = useState('')

  const shortenLink = () => {
    console.log(url)
  }
  return (
    <div className='page-container'>
        <div className="content-container">
            <div className="input-container">
                <div className="function">
                    <input 
                    type="text"
                    value={url}
                    onChange={(e) => setUrl(e.target.value)}
                    />
                    <button className="shorten-link" onClick={shortenLink}>Shrink</button>
                </div>
            </div>
            <div className="data-container">

            </div>
        </div>
    </div>
  )
}

export default AdminPage