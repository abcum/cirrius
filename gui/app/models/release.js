import DS from 'ember-data';
import Fixtures from "gui/fixtures/release";

export default DS.Model.extend({

	project: DS.belongsTo('project'),

	time: DS.attr('date', {
    	defaultValue: function() {
        	return new Date();
        }
    }),

    name: DS.attr('string', {
		defaultValue: ''
	}),

	info: DS.attr('string', {
		defaultValue: ''
	}),

}).reopenClass(Fixtures);
