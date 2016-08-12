import DS from 'ember-data';
import Fixtures from "gui/fixtures/file";

export default DS.Model.extend({

	project: DS.belongsTo('project', {
        async: true,
    }),
	
	path: DS.attr('string', {
        defaultValue: ''
    }),

    content: DS.attr('string', {
        defaultValue: ''
    }),

    created_at: DS.attr('date', {
        defaultValue: function() {
        	return new Date();
        }
    }),

    updated_at: DS.attr('date', {
        defaultValue: function() {
        	return new Date();
        }
    }),

    summary: Ember.computed('path', function() {
		return this.get('path') || '<span style="color:red">No name</span>'.risk();
	}),

}).reopenClass(Fixtures);