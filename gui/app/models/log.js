import DS from 'ember-data';
import Fixtures from "gui/fixtures/log";

export default DS.Model.extend({

	type: DS.attr('string', {
        defaultValue: ''
    }),

    fold: DS.attr('string', {
        defaultValue: ''
    }),

    file: DS.attr('string', {
        defaultValue: ''
    }),

    line: DS.attr('string', {
        defaultValue: ''
    }),

 	char: DS.attr('string', {
        defaultValue: ''
    }),

    args: DS.attr(),

}).reopenClass(Fixtures);
