import Ember from 'ember';

export default Ember.Route.extend({

	model(params, transition) {
		console.log(transition.params.project.project_id);
		return this.store.query('request', {
			filter: { project: transition.params.project.project_id }
		});
	}

});
