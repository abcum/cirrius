import Ember from "ember";

export default Ember.Controller.extend({

    actions: {

        create: function() {

			let name = this.get('newdomain');

			if (Ember.isEmpty(name)) return;

            this.store.createRecord('domain', {
            	name: name,
            	project: this.get('model')
            }).save().then( () => {
            	this.set('newdomain', '');
            });

        },

    }

});
