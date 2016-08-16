import Ember from "ember";

export default Ember.Controller.extend({

    file: Ember.inject.controller(),
    request: Ember.inject.controller(),

    // Ember.RouteMixin, Ember.RmenuMixin, Ember.SelectableMixin,

    // search: '',

    // searchProperties: ['name:or'],

    // searchedContent: Ember.computed.search('model', 'search', 'searchProperties'),

    // sortProperties: ['name:asc'],

    // arrangedContent: Ember.computed.sort('searchedContent', 'sortProperties'),

    // selectable: Ember.computed.alias('arrangedContent'),

    actions: {

        create: function() {

            this.store
                .createRecord('file', { project: this.get('model') })
                .save()
                .then(model => {
                    this.transitionToRoute('file', model);
                })
            ;

        },

    }

});
