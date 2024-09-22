import { useState } from 'react';
import './css/Store.css';

function Store() {

  const [input, setInput] = useState("");
  const [imageBase64, setImageBase64] = useState("");

  const handleButtonClick = () => {
    console.log("key : "+input);
    console.log("image : "+imageBase64);

    const setImage = async () => {
      try {
        const response = await fetch('http://127.0.0.1:8080/set', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ key: input, value: imageBase64 }),
        });
        const data = await response.json();
        console.log(data);
      }catch (error) {
        console.error('Error fetching image:', error);
      }
    };
    setImage();
  };

  const handleInputChange = (event) => {
    setInput(event.target.value);

  };

  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setImageBase64(reader.result); // Set base64 string in state
      };
      reader.readAsDataURL(file);       // Convert image file to base64
    }
  };

  return (
    <>
      <div className='store-div'>
        <input
          type="text"
          id="input"
          value={input}
          placeholder="Enter key for image"
          onChange={handleInputChange} 
        />

        <input
          type="file"
          accept="image/*"
          onChange={handleFileChange}
        />

        <button onClick={handleButtonClick}>
          Submit
        </button>
      </div>
    </>
  );
}

export default Store;
