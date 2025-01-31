export let cart = JSON.parse(localStorage.getItem('cart')) || [{
    productId: "e43638ce-6aa0-4b85-b27f-e1d07eb678c6",
    quantity: 1
}];

function saveToStorage() {
    localStorage.setItem('cart', JSON.stringify(cart));
}

export function updateDeliveryOption(productId, deliveryOptionId) {
    let matchingItem;
  
    cart.forEach((cartItem) => {
      if (productId === cartItem.productId) {
        matchingItem = cartItem;
      }
    });
  
    matchingItem.deliveryOptionId = deliveryOptionId;
  
    saveToStorage();
  }

export function addToCart(productId) {
    let matchingItem;

    cart.forEach((cartItem) => {
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

    saveToStorage();
}

export function removeFromCart(productId) {
    const new_cart = [];
    cart.forEach((cartItem) => {
        if (cartItem.productId !== productId) {
            new_cart.push(cartItem);
        }
    });

    cart = new_cart;

    saveToStorage(); 
}

export function updateQuantity(productId, newQuantity) {
    cart.forEach((cartItem) => {
        if (cartItem.productId === productId) {
            cartItem.quantity = newQuantity;
        }
    })

    saveToStorage();
}

export function calculateCartQuantity() {
    let cartQuantity = 0;

    cart.forEach((cartItem) => {
        cartQuantity += cartItem.quantity;
    });

    return cartQuantity;
}