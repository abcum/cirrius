import Ember from 'ember';

export default Ember.Component.extend({

	tagName: 'tree-view',

	didReceiveAttrs() {

		let path = this.get('path');

		Ember.defineProperty(this, 'content', Ember.computed('files', `files.@each.${path}`, 'value', function() {
			
		    return this.get('files').reduce( (prev, item) => {
		        let dirs = item.get(path	).split('/');
		        let file = dirs.pop();
		        let tree = dirs.reduce( (h, d) => h.fo[d] = h.fo[d] || { fo: {}, fi: {} }, prev);
		        return (tree.fi[file] = item), prev;
		    }, { fo: {}, fi: {} });

		}));

	},

});
