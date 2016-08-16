import Ember from "ember";

export default Ember.Controller.extend({

    // Ember.RouteMixin, Ember.RmenuMixin, Ember.SelectableMixin,

    // search: '',

    // searchProperties: ['name:or'],

    // searchedContent: Ember.computed.search('model', 'search', 'searchProperties'),

    // sortProperties: ['name:asc'],

    // arrangedContent: Ember.computed.sort('searchedContent', 'sortProperties'),

    // selectable: Ember.computed.alias('arrangedContent'),

    actions: {

        delete: function() {

            this.model.destroyRecord().then( () => {
                this.transitionToRoute('project.index');
            });

        },

        /*setupController: function(controller, model) {

            this._super(controller, model);

            this.controllerFor('files').send('select', model);

         },*/

    }

});
