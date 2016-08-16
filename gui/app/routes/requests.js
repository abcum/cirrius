import Ember from 'ember';

export default Ember.Route.extend({

	model(params, transition) {
		return this.store.query('request', {
			filter: { project: transition.params.project.project_id }
		});
	}

});
