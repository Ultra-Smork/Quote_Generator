import { useEffect, useState } from 'react'
import './App.css'
const API_URL = "http://localhost:8080/CHEBUREK"

function App() {
  const [quote, setQuote] = useState("")
  const [author, setAuthor] = useState("")
    const FetchQuote = async () =>{
        const response = await fetch(API_URL)
        const json_reponse = await response.json()
        setQuote(json_reponse[0].quote)
        setAuthor(json_reponse[0].author)
        console.log(json_reponse)
    }
  useEffect(() =>{
    FetchQuote()
  }, [])

  return (
    <>
       <div className='Title'>Quote Generator</div>
       <div className='Card'>
          <div className='Screen Title Screen2'>{author}</div>
         <div className='Screen Quote'>{quote}</div>
       </div>
       <div className='GenDis'>
          <button className='Button1' onClick={() => FetchQuote()}>Generate quote</button>
       </div>
    </>
  )
}

export default App;
