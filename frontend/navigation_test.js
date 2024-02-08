Feature('menu');

Scenario('should display menu products', ({ I }) => {
  I.amOnPage('/');
  I.see('FULL HOUSE');
  I.see('RESTOBAR');


  I.seeElement('.product-card'); // Verificar que al menos un producto se muestra
  I.seeNumberOfElements('.product-card', 5); // Verificar que hay 4 productos
  
}).config({ "waitForTimeout": 50000 }); 
