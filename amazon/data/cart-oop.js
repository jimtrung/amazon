const cart = {
  cartItems: JSON.parse(localStorage.getItem('cart-oop')) || undefined,

  updateDeliveryOption(productId, deliveryOptionId) {
    let matchingItem;
  
    this.cartItems.forEach((cartItem) => {
      if (productId === cartItem.productId) {
        matchingItem = cartItem;
      }
    });
  
    matchingItem.deliveryOptionId = deliveryOptionId;
  
    this.saveToStorage();
  },

  saveToStorage() {
    localStorage.setItem('cart-oop', JSON.stringify(this.cartItems));
  },

  updateDeliveryOption(productId, deliveryOptionId) {
    let matchingItem;
  
    this.cartItems.forEach((cartItem) => {
      if (productId === cartItem.productId) {
        matchingItem = cartItem;
      }
    });
    matchingItem.deliveryOptionId = deliveryOptionId;
    this.saveToStorage();
  },

  addToCart(productId) {
    let matchingItem;

    this.cartItems.forEach((cartItem) => {
        if (cartItem.productId === productId) {
            matchingItem = cartItem;
        }
    });

    if (matchingItem) {
        matchingItem.quantity += Number(
        document.querySelector(`.js-quantity-selector-${productId}`).value
        ); 
    }
    else {
        cart.push({
            productId,
            quantity: Number(document.querySelector(`.js-quantity-selector-${productId}`).value),
            deliveryOptionId: '1'
        });
    }
    this.saveToStorage();
  },

  removeFromCart(productId) {
    const new_cart = [];
    this.cartItems.forEach((cartItem) => {
        if (cartItem.productId !== productId) {
            new_cart.push(cartItem);
        }
    });
    this.cartItems = new_cart;
    this.saveToStorage(); 
  },

  updateQuantity(productId, newQuantity) {
    this.cartItems.forEach((cartItem) => {
      if (cartItem.productId === productId) {
        cartItem.quantity = newQuantity;
      }
    })
    this.saveToStorage();
  },

  calculateCartQuantity() {
    let cartQuantity = 0;

    this.cartItems.forEach((cartItem) => {
        cartQuantity += cartItem.quantity;
    });

    return cartQuantity;
  }
};

console.log(cart);