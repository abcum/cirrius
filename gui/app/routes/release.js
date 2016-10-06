import Ember from 'ember';

export default Ember.Route.extend({
	deactivate() {
		this.controllerFor('release').set('model', null);
	}
});
