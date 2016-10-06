import Ember from 'ember';

export default Ember.Controller.extend({

	editing: false,

	actions: {

		live: function() {
            this.get('model').set('live', true);
			this.get('model').set('test', false);
		},

		test: function() {
            this.get('model').set('test', true);
			this.get('model').set('live', false);
		},

		delete: function() {

			let model = this.get('model.live');

			if (model==true) {
				alert('Cannot delete the current live release');
				return
			};

			this.get('model').destroyRecord()
				.then( () => {
					this.transitionToRoute('releases');
				})
		},
	}
});
