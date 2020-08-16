paypal.Buttons({

    // Set your environment
    env: paypal_env,

    // Set style of buttons
    style: {
        layout: 'horizontal',   // horizontal | vertical
        size: 'large',   // medium | large | responsive
        shape: 'rect',         // pill | rect
        color: 'gold',         // gold | blue | silver | black,
        fundingicons: false,    // true | false,
        tagline: false          // true | false,
    },
// onInit is called when the button first renders
    onInit: function (data, actions) {

        // Disable the buttons
        actions.disable();

        // Listen for changes to the checkbox
        document.querySelector('#amount')
            .addEventListener('keyup', function (event) {
                let amount = event.target.value;
                amount = +amount;
                if (isNaN(amount)) {
                    document.getElementById("amount-err").innerHTML = "Invalid amount format";
                    setTimeout(function () {
                        document.getElementById("amount-err").innerHTML = "";
                    }, 2000);
                    actions.disable();
                } else if (amount < 5) {
                    document.getElementById("amount-err").innerHTML = "Amount must be greater than 5";
                    setTimeout(function () {
                        document.getElementById("amount-err").innerHTML = "";
                    }, 2000);
                    actions.disable();
                }
                if (event.target.value === "") {
                    document.getElementById("amount-err").innerHTML = "Invalid amount format";
                    setTimeout(function () {
                        document.getElementById("amount-err").innerHTML = "";
                    }, 2000);
                    actions.disable();
                } else {
                    actions.enable();
                }
            });
    },

    // onClick is called when the button is clicked
    onClick: function () {
        var regex = /^\d+(?:\.\d{0,2})$/;
        let amount = document.getElementById("amount").value;
        amount = +amount;
        if (isNaN(amount)) {
            alert(2);
            document.getElementById("amount-err").innerHTML = "Invalid amount format";
            setTimeout(function () {
                document.getElementById("amount-err").innerHTML = "";
            }, 2000);
        }
    },
    // Wait for the PayPal button to be clicked
    createOrder: function () {
        let formData = new FormData();
        let amount = document.getElementById("amount").value;

        formData.append('amount', amount);

        return fetch(
            '/account/paypal/do/order',
            {
                method: 'POST',
                body: formData
            }
        ).then(function (response) {
            return response.json();
        }).then(function (resJson) {
            return resJson.data.id;
        });
    },

    // Wait for the payment to be authorized by the customer
    onApprove: function (data, actions) {
        return fetch(
            '/account/paypal/success/' + data.orderID,
            {
                method: 'POST'
            }
        ).then(function (res) {
            return res.json();
        }).then(function (res) {

            // window.location.href = '/account/paypal/order/success';
        });
    },


    // Wait for the payment to be authorized by the customer
    onCancel: function (data, actions) {
        return fetch(
            '/account/paypal/cancel/' + data.orderID,
            {
                method: 'POST'
            }
        ).then(function (res) {
            return res.json();
        }).then(function (res) {

            // window.location.href = '/account/paypal/order/success';
        });
    },


}).render('#paypalCheckoutContainer');