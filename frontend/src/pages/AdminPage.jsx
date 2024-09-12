import React, { useState } from 'react'

function AdminPage() {
  const [url, setUrl] = useState('')

  const shortenLink = (e) => {
    e.preventDefault()
    if(isUrlValid(url)){

    } else {
        
    }
  }

  function isUrlValid(input) {
    var res = input.match(/(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g);
    if(res == null){
        return false;
    } else {
        return true;
    }
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