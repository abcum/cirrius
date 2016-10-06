import DS from 'ember-data';
import Fixtures from "gui/fixtures/stage";

export default DS.Model.extend({

	project: DS.belongsTo('project'),

	type: DS.attr('string', {
		defaultValue: ''
	}),

}).reopenClass(Fixtures);
