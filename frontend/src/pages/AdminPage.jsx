import React, { useState } from 'react'

function AdminPage() {
  const [url, setUrl] = useState('')

  const shortenLink = (e) => {
    e.preventDefault()
    
  }

  

  return (
    <div className='page-container'>
        <div className="content-container">
            <div className="input-container">
                <div className="function">
                    <form onSubmit={shortenLink}>
                        <input 
                        type="text"
                        value={url}
                        onChange={(e) => setUrl(e.target.value)}
                        required
                        />

                        <button className="shorten-link" type='submit'>Shrink</button>
                    </form>               
                </div>
            </div>
            <div className="data-container">

            </div>
        </div>
    </div>
  )
}

export default AdminPage