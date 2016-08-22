import DS from 'ember-data';
import Fixtures from "gui/fixtures/request";

export default DS.Model.extend({

	project: DS.belongsTo('project'),

	status: DS.attr('number', {
		defaultValue: ''
	}),

	url: DS.attr('string', {
		defaultValue: ''
	}),

	duration: DS.attr('string', {
		defaultValue: ''
	}),

	ip: DS.attr('string', {
		defaultValue: ''
	}),

	method: DS.attr('string', {
		defaultValue: ''
	}),

	size: DS.attr('number', {
		defaultValue: ''
	}),

	time: DS.attr('date', {
		defaultValue: function() {
			return new Date();
		}
	}),

	logs: DS.hasMany('log', {
		async: true,
	}),

}).reopenClass(Fixtures);
