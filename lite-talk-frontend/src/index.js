import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import AppExample from './App';
import Header from './components/header';
import './index.css';

ReactDOM.render(
  <React.StrictMode>
    <Header />
    <AppExample />
  </React.StrictMode>,
  document.getElementById('root')
);
