import { cart, removeFromCart, updateQuantity, updateDeliveryOption }
    from '../../data/cart.js';
import { getProduct } from '../../data/products.js';
import { formatCurrency } from '../utils/money.js';
import { deliveryOptions, getDeliveryOption, calculateDeliveryDate } 
    from '../../data/deliveryOption.js';
import { renderPaymentSummary } from './paymentSummary.js';
import { renderCheckoutHeader } from './checkoutHeader.js';


export function renderOrderSummary() {
    let cartSummaryHTML = '';

    cart.forEach((cartItem) => {
        const productId = cartItem.productId;
        const matchingProduct = getProduct(productId);

        const deliveryOptionId = cartItem.deliveryOptionId;
        const deliveryOption = getDeliveryOption(deliveryOptionId);
        const dateString = calculateDeliveryDate(deliveryOption);

        cartSummaryHTML += `
        <div class=" cart-item-container js-cart-item-container-${matchingProduct.id}">
            <div class="delivery-date">
                Delivery date: ${dateString}
            </div>

            <div class="cart-item-details-grid">
                <img class="product-image"
                    src="${matchingProduct.image}">

                <div class="cart-item-details">
                    <div class="product-name">
                        ${matchingProduct.name}
                    </div>
                    <div class="product-price">
                        $${formatCurrency(matchingProduct.priceCents)}
                    </div>
                    <div class="product-quantity">
                        <span>
                            Quantity: <span class="quantity-label js-quantity-label-${matchingProduct.id}">
                            ${cartItem.quantity}</span>
                        </span>
                        <span class="update-quantity-link link-primary js-update-link"
                        data-product-id=${matchingProduct.id}>
                            Update
                        </span>
                        <input type="number" class="quantity-input-${productId} js-quantity-input is-invisible"
                        data-product-id=${matchingProduct.id} value="${cartItem.quantity}">
                        <span class="save-quantity-link-${productId} link-primary js-save-link is-invisible"
                        data-product-id=${matchingProduct.id}>Save</span>
                        <span class="delete-quantity-link link-primary js-delete-link" 
                        data-product-id=${matchingProduct.id}>
                            Delete
                        </span>
                    </div>
                </div>

                <div class="delivery-options">
                    <div class="delivery-options-title">
                        Choose a delivery option:
                    </div>
                    ${deliveryOptionsHTML(matchingProduct, cartItem)}
                </div>
            </div>
        </div>
        `
    });
    document.querySelector('.js-order-summary').innerHTML = cartSummaryHTML;

    function deliveryOptionsHTML(matchingProduct, cartItem) {
        let HTML = ''

        deliveryOptions.forEach((deliveryOption) => {
            const dateString = calculateDeliveryDate(deliveryOption);
            const priceString = deliveryOption.priceCents === 0 ? 'FREE'
                : `$${formatCurrency(deliveryOption.priceCents)}`;

            const isChecked = deliveryOption.id === cartItem.deliveryOptionId;

            HTML += `
            <div class="delivery-options js-delivery-options"
            data-product-id="${matchingProduct.id}"
            data-delivery-option-id="${deliveryOption.id}">
                <div class="delivery-option">
                    <input type="radio" ${isChecked ? 'checked' : ''}
                        class="delivery-option-input"
                        name="delivery-option-${matchingProduct.id}">
                    <div>
                        <div class="delivery-option-date">
                            ${dateString}
                        </div>
                        <div class="delivery-option-price">
                            ${priceString} Shipping
                        </div>
                    </div>
                </div>
            </div>
            `
        });

        return HTML;
    }

    document.querySelectorAll('.js-delete-link').forEach((link) => {
        link.addEventListener('click', () => {
            const productId = link.dataset.productId;
            removeFromCart(productId);

            renderPaymentSummary();
            renderOrderSummary();
            renderCheckoutHeader();
        });
    });

    document.querySelectorAll('.js-update-link').forEach((link) => {
        link.addEventListener('click', () => {
            const productId = link.dataset.productId;

            const quantityInput = document.querySelector(`.quantity-input-${productId}`);
            const saveButton = document.querySelector(`.save-quantity-link-${productId}`);

            quantityInput.classList.remove('is-invisible');
            saveButton.classList.remove('is-invisible');

            link.classList.add('is-invisible');

            document.querySelector(`.js-quantity-label-${productId}`).innerHTML = '';
        });
    });

    document.querySelectorAll('.js-save-link').forEach((link) => {
        link.addEventListener('click', () => {
            const productId = link.dataset.productId;
            const quantityInput = document.querySelector(`.quantity-input-${productId}`);
            const updateButton = document.querySelector('.js-update-link');

            updateButton.classList.remove('is-invisible');
            quantityInput.classList.add('is-invisible');
            link.classList.add('is-invisible');

            const newQuantity = Number(quantityInput.value);
            if (newQuantity === 0) {
                removeFromCart(productId);
                const container = document.querySelector(`.js-cart-item-container-${productId}`);

                container.remove();
            }
            else if (newQuantity < 0) {
                alert('Not a valid quantity.');
            }
            else {
                updateQuantity(productId, newQuantity);

                document.querySelector(`.js-quantity-label-${productId}`).innerHTML = newQuantity;
                renderPaymentSummary();
                renderOrderSummary();
                renderCheckoutHeader();    
            }
        });
    });

    document.querySelectorAll('.js-delivery-options')
        .forEach((element) => {
            element.addEventListener('click', () => {
                const { productId, deliveryOptionId } = element.dataset;
                updateDeliveryOption(productId, deliveryOptionId);
                renderOrderSummary();        
                renderPaymentSummary();        
            });
        });
}