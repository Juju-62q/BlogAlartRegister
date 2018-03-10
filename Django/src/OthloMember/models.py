from django.db import models

# Create your models here.
class Member(models.Model):
    """メンバー"""
    name = models.CharField('name', max_length=255)
    slackName = models.CharField('slack_name', max_length=255)
    graduateDate = models.DateTimeField('graduate_date')

    def __str__(self):
        return self.name