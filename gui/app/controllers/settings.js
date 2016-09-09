import Ember from "ember";

export default Ember.Controller.extend({

    actions: {

        createDomain: function() {

			let name = this.get('newdomain');

			if (Ember.isEmpty(name)) return;

            this.store.createRecord('domain', {
            	name: name,
            	project: this.get('model')
            }).save().then( () => {
            	this.set('newdomain', '');
            });

        },

        deleteDomain: function(domain) {
            domain.destroyRecord();
        },

        createStage: function() {

			let type = this.get('newstage');

			if (Ember.isEmpty(type)) return;

            this.store.createRecord('stage', {
            	type: type,
            	project: this.get('model')
            }).save().then( () => {
            	this.set('newstage', '');
            });

        },

        deleteStage: function(stage) {
            stage.destroyRecord();
        },

    }

});
