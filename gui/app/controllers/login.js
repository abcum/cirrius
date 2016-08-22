import Ember from 'ember';

export default Ember.Controller.extend({

	pass: null,

	loading: false,

    actions: {

        submit: function() {
        	console.log(this.get('storage.email') + ':' + this.get('pass'));
        }

    },
});
