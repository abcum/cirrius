import Ember from "ember";

export default Ember.Controller.extend({

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