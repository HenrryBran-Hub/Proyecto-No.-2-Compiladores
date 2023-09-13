import React, { useRef, useState } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

const App = () => {
  const fileInputRef = useRef(null);
  const textarea1Ref = useRef(null);
  const [code, setCode] = useState('');
  const [currentFile, setCurrentFile] = useState(null);
  const [result, setResult] = useState('');

  const handleAbrirClick = () => {
    fileInputRef.current.click();
  };

  const handleFileSelect = (event) => {
    const file = event.target.files[0];
    if (file) {
      setCurrentFile(file);
      const reader = new FileReader();
      reader.onload = (e) => {
        textarea1Ref.current.value = e.target.result;
      };
      reader.readAsText(file);
    }
  };

  const handleGuardarClick = () => {
    if (currentFile) {
      const blob = new Blob([textarea1Ref.current.value], {
        type: 'text/plain',
      });
      const url = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = currentFile.name;
      a.click();
      URL.revokeObjectURL(url);
      alert('Archivo actualizado');
    }
  };

  const handleGuardarComoClick = () => {
    const blob = new Blob([textarea1Ref.current.value], { type: 'text/plain' });
    const url = URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'archivo.txt';
    a.click();
    URL.revokeObjectURL(url);
  };

  const handleAcercaDeClick = () => {
    alert('Henrry David Bran Velasquez \n201314439');
  };

  const handleEjecutarClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/ejecutar', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: textarea1Ref.current.value }),
      });

      if (response.ok) {
        const data = await response.json();
        setResult(data.result);
        alert('Compilado con Exito');
      } else {
        alert('Error en la respuesta:', response.statusText);
      }
    } catch (error) {
      alert('Error en la solicitud:', error);
    }
  };

  const handleCstClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/arbol', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: textarea1Ref.current.value }),
      });

      if (response.ok) {
        const data = await response.json();
        setResult(data.result);
        alert('Compilado con Exito');
      } else {
        alert('Error en la respuesta:', response.statusText);
      }
    } catch (error) {
      alert('Error en la solicitud:', error);
    }
  };

  const handleSimbolosClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/simbolos', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: textarea1Ref.current.value }),
      });

      if (response.ok) {
        const data = await response.json();
        setResult(data.result);
        alert('Compilado con Exito');
      } else {
        alert('Error en la respuesta:', response.statusText);
      }
    } catch (error) {
      alert('Error en la solicitud:', error);
    }
  };

  const handleErroresClick = async () => {
    try {
      const response = await fetch('http://localhost:8080/errores', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code: textarea1Ref.current.value }),
      });

      if (response.ok) {
        const data = await response.json();
        setResult(data.result);
        alert('Compilado con Exito');
      } else {
        alert('Error en la respuesta:', response.statusText);
      }
    } catch (error) {
      alert('Error en la solicitud:', error);
    }
  };

  return (
    <>
      <input
        type="file"
        ref={fileInputRef}
        onChange={handleFileSelect}
        style={{ display: 'none' }}
      />
      <nav className="navbar navbar-dark bg-dark">
        <span className="navbar-brand mb-0 h1 text-light">T-Swift IDE</span>
      </nav>
      <div className="container-fluid bg-dark text-light">
      <div className="col-md-12 text-left">
            <button id="abrir" className="btn btn-outline-primary" onClick={handleAbrirClick}>Abrir</button>
            <button id="guardar" className="btn btn-outline-primary" onClick={handleGuardarClick}>Guardar</button>
            <button id="guardar-como" className="btn btn-outline-primary" onClick={handleGuardarComoClick}>GuardarComo</button>
            <button id="acerca-de" className="btn btn-outline-primary" onClick={handleAcercaDeClick}>Acerca de</button>
        </div>
      </div>
      <div className="container-fluid bg-dark text-light">
        <div className="row">
          <div className="col-md-6">
            <label htmlFor="textarea1">Editor</label>
            <hr />
            <textarea
              id="textarea1"
              ref={textarea1Ref}
              className="form-control bg-secondary text-light"
              rows="22"
              style={{ resize: 'none' }}
              value={code}
              onChange={(e) => setCode(e.target.value)}
            ></textarea>
          </div>
          <div className="col-md-6">
            <label htmlFor="textarea2">Consola</label>
            <hr />
            <textarea
              id="textarea2"
              className="form-control bg-secondary text-light"
              rows="22"
              style={{ resize: 'none' }}
              readOnly
              value={result}
            ></textarea>
          </div>
        </div>
        <hr />
        <div className="row">
          <div className="col-md-12 text-center">
            <button id="ejecutar" className="btn btn-outline-success" onClick={handleEjecutarClick}>Compilar</button>
            <button id="mostrar-reportes" className="btn btn-outline-info" onClick={handleCstClick}>Optimizar</button>
            <button id="simbolos" className="btn btn-outline-warning" onClick={handleSimbolosClick}>Simbolos</button>
            <button id="errores" className="btn btn-outline-danger" onClick={handleErroresClick}>Errores</button>
          </div>
          <hr />
          <hr />
        </div>
      </div>
    </>
  );
};

export default App;
