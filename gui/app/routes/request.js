import Ember from 'ember';

export default Ember.Route.extend({
	deactivate() {
		this.controllerFor('request').set('model', null);
	}
});
