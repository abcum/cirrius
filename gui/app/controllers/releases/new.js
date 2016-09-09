import Ember from "ember";

export default Ember.Controller.extend({

	actions: {

		commit: function() {

			let name = this.get('name');
			let info = this.get('info');

			if (Ember.isEmpty(name)) return;

			this.store
				.createRecord('release', {
					name: name,
					info: info,
					project: this.get('model'),
				})
				.save()
				.then(() => {
					this.transitionToRoute('releases');
				})
			;
		},

	},
});
