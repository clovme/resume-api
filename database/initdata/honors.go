package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Honors 荣誉证书
func (d *InitData) Honors() {
	modelList := []models.Honors{
		{BaseModelWithRIDUID: ridUID(), Content: "<ol><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>项目管理专家认证（PMP）</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发机构: 项目管理协会（PMI）</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发日期: 2020年6月</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>描述: 获得PMP认证，证明具备在项目管理领域的专业知识和技能，能够有效地规划、执行和管理各类项目，确保项目按时、在预算内并符合质量标准地完成。</strong></li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>高级软件工程师认证</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发机构: 国际软件工程协会（ISEB）</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发日期: 2021年4月</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>描述: 通过严格的考试，获得高级软件工程师认证，表明在软件开发、测试、维护和项目管理方面具有深厚的专业知识和丰富的实践经验。</strong></li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>数据科学与机器学习专家认证</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发机构: 数据科学协会（DSS）</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>颁发日期: 2022年9月</strong></li><li data-list=\"ordered\" class=\"ql-indent-1\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>描述: 获得数据科学与机器学习专家认证，掌握了数据分析、统计建模、机器学习算法和大数据处理技术，能够在实际工作中应用这些技能进行数据驱动的决策和创新。</strong></li></ol>"},
	}

	insertRecord[models.Honors]("荣誉证书", modelList, func(model models.Honors) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
