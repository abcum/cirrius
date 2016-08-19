import Ember from "ember";

export default Ember.Controller.extend({

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
