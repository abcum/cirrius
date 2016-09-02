import DS from 'ember-data';
import Fixtures from "gui/fixtures/domain";

export default DS.Model.extend({

	project: DS.belongsTo('project'),

	name: DS.attr('string', {
		defaultValue: ''
	}),

}).reopenClass(Fixtures);
