import DS from 'ember-data';
import Fixtures from "gui/fixtures/project";

export default DS.Model.extend({

  	fullname: DS.attr('string', {
        defaultValue: ''
    }),

    shortname: DS.attr('string', {
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

    files: DS.hasMany('file', {
        async: true,
    }),

}).reopenClass(Fixtures);