import Ember from 'ember';

export default Ember.Component.extend({

	tagName: 'tree-view-list',

	actions: {

		file(file) {
			this.sendAction('onselect', file);
		},

	}

});
