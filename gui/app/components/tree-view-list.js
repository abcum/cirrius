import Ember from 'ember';

export default Ember.Component.extend({

	tagName: 'tree-view-list',

	actions: {

		folder(subtree) {
			Ember.set(subtree, 'show', !Ember.get(subtree, 'show'));
		},

		file(file) {
			this.sendAction('onselect', file);
		},

	}
	
});
