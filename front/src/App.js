import React, {useState} from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const [result, setResult] = useState('?');
  const [op1, setOp1] = useState(0);
  const [op2, setOp2] = useState(0);

  const handleClick = () => {
    fetch(`/sum?operand1=${op1}&operand2=${op2}`, {
      headers: {
        'Access-Control-Allow-Origin': '*',
      },
    })
      .then(response => response.json())
      .then(response => {
        setResult(response.result);
        console.log(`1 + 15 = ${response.result}`);
      });
  };

  const handleChangeOp1 = (evt) => {
    setOp1(evt.target.value);
  }

  const handleChangeOp2 = (evt) => {
    setOp2(evt.target.value);
  }

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <input type="number" value={op1} onChange={handleChangeOp1}/>
        +
        <input type="number" value={op2} onChange={handleChangeOp2} />
        <button value="=" onClick={handleClick}>=</button>
        {result}
      </header>
    </div>
  );
}

export default App;
