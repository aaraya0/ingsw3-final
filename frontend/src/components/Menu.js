import React, { useState, useEffect } from 'react';
import '../styles/Menu.css';

function Menu() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await fetch('http://localhost:8090/products');
        const data = await response.json();
        setProducts(data);
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    };

    fetchProducts();
  }, []);

  return (

    <div className="menu">
         
       <div className="left-side"></div> <div className="title"><h1 className="page-title neonText">FULL HOUSE</h1></div>  <div className="right-side"></div>
       <div className="subtitle"> <h2>RESTOBAR</h2></div>
        <div className="background"><video
        autoPlay
        loop
        muted
        className="background-video"
      >
        <source src="../images/background-video.mp4" type="video/mp4"/>
        Tu navegador no soporta el elemento de video.
      </video>
      </div>
         <div className="menu-container">
      {products.map((product) => (
        <div key={product.id} className="product-card">
          <img
            src={`images/${product.image}`}
            alt={product.title}
            className="product-image"
          />
          <h3 className="product-title">{product.title}</h3>
          <p className="product-price">Precio: ${product.price}</p>
          <p className="product-description">{product.description}</p>
        </div>
      ))}
      </div>
      <footer>
      <p>Autor: Agustina Araya</p>
    </footer>
    </div>
  );
}

export default Menu;
