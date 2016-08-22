import Ember from "ember";

export default Ember.Route.extend({

	model(params) {
	    return this.store.findRecord('file', params.file_id);
	},

	deactivate() {
		this.controllerFor('file').set('model', null);
	}

	/*setupController(controller, model) {
		this._super(...arguments);
		this.controllerFor('project').set('file', model);
	},

	deactivate() {
		this._super(...arguments);
		this.controllerFor('project').set('file', null);
	}*/

});