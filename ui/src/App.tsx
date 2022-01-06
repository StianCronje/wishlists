import React from 'react';
import './App.css';
import { ItemList } from './components/ItemList';

function App() {
  return (
    <div className="App">
      <ItemList items={[]} />
    </div>
  );
}

export default App;
