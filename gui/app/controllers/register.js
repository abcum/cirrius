import Ember from "ember";

export default Ember.Controller.extend({

    error: null,

    loading: false,

    complete: false,

    actions: {

        submit: function() {

            var self = this;

            // Mark the request as loading
            // so that button is greyed-out

            this.set('loading', true);

            // Make the request to the api
            // and wait for the response

            // TODO Need to implement surreal registration
        }

    }

});