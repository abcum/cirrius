import Ember from "ember";

export default Ember.Controller.extend({

	actions: {

		create: function() {

			let name = window.prompt("Project name:");

			if (Ember.isEmpty(name)) return;

			this.store
				.createRecord('project', { name: name })
				.save()
				.then(model => {
					this.transitionToRoute('project', model);
				})
			;

		},

	}

});
