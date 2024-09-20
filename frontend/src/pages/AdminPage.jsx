import React, { useState } from 'react'
import DataView from '../components/DataView'
import '../pages/AdminPage.css'

function AdminPage() {
  const [url, setUrl] = useState('')
  const [error, setError] = useState('')


  const shortenLink = (e) => {
    e.preventDefault()
    if(isUrlValid(url)){
        shrinkUrl()
        setError('')
    } else {
        //display an error message to the user
        setError("Please enter a valid url")
    }
  }

  const shrinkUrl = async () => {
    try {
      const response = await axios.post('http://localhost:1323/shrink', {
        url: url, 
      });
  
      setData(response.data); 
    } catch (error) {
      console.error('Error shortening URL:', error); 
    }
  };
  

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
                <DataView/>
            </div>
        </div>
    </div>
  )
}

export default AdminPage