import Ember from "ember";

export default Ember.Controller.extend({

    actions: {

        create: function() {

            this.store
                .createRecord('project')
                .save()
                .then(model => {
                    this.transitionToRoute('project', model);
                })
            ;

        },

    }

});
