import React, { useState } from 'react';
import './css/Retrieve.css';

function Retrieve() {
  const [input, setInput] = useState("");
  const [imageBase64, setImageBase64] = useState(""); 

  const handleInputChange = (event) => {
    setInput(event.target.value);
  };

  const handleSubmit = async () => {
    getImage(); // Call the function to retrieve the image
  };

  // api to retrieve base64 string using key from backend
  const getImage = async () => {
    try {
      const response = await fetch(`http://127.0.0.1:8080/get?key=${input}`);
      const data = await response.json();
      if(data.value=="Does not exist"){
        console.log("Does not exist");
        setImageBase64("");
      }else{
        setImageBase64(data.value);       // Set the base64 string from the response
        console.log(data.value);
    }
    } catch (error) {
      console.error('Error fetching image:', error);
    }
  };

  return (
    <div>
      <input
        type="text"
        id="input"
        placeholder="Enter key"
        value={input}
        onChange={handleInputChange}
      />
      <button onClick={handleSubmit}>Submit</button>

      {imageBase64 && (
        <div>
          <h3>Image Preview:</h3>
          <img src={imageBase64} alt="Uploaded" style={{ maxWidth: '100%', height: 'auto' }} />
        </div>
      )}
    </div>
  );
}

export default Retrieve;
