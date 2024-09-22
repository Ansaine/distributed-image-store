import { useState } from 'react'
import './App.css'
import Store from './components/store'
import Retrieve from './components/Retrieve'

function App() {

  return (
    <>
      <h2>Distributed Image Store</h2>
      <Store/>
      <Retrieve/>
    </>
  )
}

export default App
